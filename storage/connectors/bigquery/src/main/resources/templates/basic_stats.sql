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
SELECT
    "{{ field.name }}" as field_name,
    -- total count
    COUNT(*) AS total_count,
    -- count
    COUNT({{ field.name }}) as feature_count,
    -- missing
    COUNT(*) - COUNT({{ field.name }}) as missing_count,
    {% if field.type equals "NUMERIC" %}
    -- mean
    AVG({{ field.name }}) as mean,
    -- stdev
    STDDEV({{ field.name }}) as stdev,
    -- zeroes
    COUNTIF({{ field.name }} = 0) as zeroes,
    -- min
    MIN({{ field.name }}) as min,
    -- max
    MAX({{ field.name }}) as max,
    -- hist will have to be called separately
    -- quantiles
    APPROX_QUANTILES(CAST({{ field.name }} AS FLOAT64), 10) AS quantiles,
    -- unique
    null as unique
    {% elseif field.type equals "CATEGORICAL" %}
    -- mean
    null as mean,
    -- stdev
    null as stdev,
    -- zeroes
    null as zeroes,
    -- min
    null as min,
    -- max
    null as max,
    -- quantiles
    ARRAY<FLOAT64>[] AS quantiles,
    -- unique
    COUNT(DISTINCT({{ field.name }})) as unique
    {% elseif field.type equals "BYTES" %}
    -- mean
    AVG(BIT_COUNT({{ field.name }})) as mean,
    -- stdev
    null as stdev,
    -- zeroes
    null as zeroes,
    -- min
    MIN(BIT_COUNT({{ field.name }})) as min,
    -- max
    MAX(BIT_COUNT({{ field.name }})) as max,
    -- hist will have to be called separately
    -- quantiles
    ARRAY<FLOAT64>[] AS quantiles,
    -- unique
    COUNT(DISTINCT({{ field.name }})) as unique
    {% elseif field.type equals "LIST" %}
    -- mean
    AVG(ARRAY_LENGTH({{ field.name }})) as mean,
    -- stdev
    null as stdev,
    -- zeroes
    null as zeroes,
    -- min
    MIN(ARRAY_LENGTH({{ field.name }})) as min,
    -- max
    MAX(ARRAY_LENGTH({{ field.name }})) as max,
    -- hist will have to be called separately
    -- quantiles
    ARRAY<FLOAT64>[] AS quantiles,
    -- unique
    null as unique
    {% endif %}
FROM subset
{% if loop.last %}{% else %}UNION ALL {% endif %}
{% endfor %}

