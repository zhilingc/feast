WITH subset AS (
    SELECT rw.* FROM (
        SELECT ARRAY_AGG(
          t ORDER BY t.created_timestamp DESC LIMIT 1
        )[OFFSET(0)]  rw
        FROM `{{ projectId }}.{{ datasetId }}.{{ featureSet.project }}_{{ featureSet.name }}_v{{ featureSet.version }}` t
        {% if featureSet.datasetId == "" %}
        WHERE event_timestamp >= '{{ featureSet.date }} 00:00:00 UTC' AND event_timestamp < DATETIME_ADD('{{ featureSet.date }}  00:00:00 UTC', INTERVAL 1 DAY)
        {% else %}
        WHERE dataset_id='{{ featureSet.datasetId }}'
        {% endif %}
        GROUP BY dataset_id, event_timestamp, {{ featureSet.entityNames | join(', ')}}
    )
)
{% for field in featureSet.fields %}
, {{ field.name }}_stats AS (
{% if field.type == 'NUMERIC' %}
  WITH stats AS (
    SELECT min+step*i as min, min+step*(i+1) as max
    FROM (
      SELECT MIN({{ field.name }}) as min, MAX({{ field.name }}) as max, (MAX({{ field.name }})-MIN({{ field.name }}))/10 step, GENERATE_ARRAY(0, 10, 1) i
      FROM subset
    ), UNNEST(i) i
  ), counts as (
    SELECT COUNT(*) as count, min, max,
    FROM subset
    JOIN stats
    ON subset.{{ field.name }} >= stats.min AND subset.{{ field.name }}<stats.max
    GROUP BY min, max
  )
  SELECT '{{ field.name }}' as field, ARRAY_AGG(STRUCT(count as count, min as low_value, max as high_value)) as num_hist, ARRAY<STRUCT<value STRING, count INT64>>[] as cat_hist FROM counts
{% elseif field.type == 'CATEGORICAL' %}
  WITH counts AS (
    SELECT {{ field.name }}, COUNT({{ field.name }}) AS count FROM subset GROUP BY {{ field.name }}
  )
  SELECT '{{ field.name }}' as field, ARRAY<STRUCT<count INT64, low_value FLOAT64, high_value FLOAT64>>[] as num_hist, ARRAY_AGG(STRUCT({{ field.name }} as value, count as count)) as cat_hist FROM counts
{% elseif field.type == 'BYTES' %}
  SELECT '{{ field.name }}' as field, ARRAY<STRUCT<count INT64, low_value FLOAT64, high_value FLOAT64>>[] as num_hist, ARRAY<STRUCT<value STRING, count INT64>>[] as cat_hist
{% elseif field.type == 'LIST' %}
  SELECT '{{ field.name }}' as field, ARRAY<STRUCT<count INT64, low_value FLOAT64, high_value FLOAT64>>[] as num_hist, ARRAY<STRUCT<value STRING, count INT64>>[] as cat_hist
{% endif %}
)
{% endfor %}
{% for field in featureSet.field %}
SELECT * FROM {{ field.name }}_stats
{% if loop.last %}{% else %}UNION ALL {% endif %}
{% endfor %}