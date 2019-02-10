resource "google_compute_instance" "redis_vm" {
  name         = "${var.name}"
  machine_type = "${var.type}"
  zone         = "${var.zone}"

  tags = ["${var.internet_tag}"]

  boot_disk {
    initialize_params {
      size  = "${var.boot_disk_size}"
      type  = "${var.boot_disk_type}"
      image = "${var.boot_disk_image}"
    }
  }

  lifecycle {
    ignore_changes = ["attached_disk"]
  }

  scheduling {
    automatic_restart = true
  }

  network_interface {
    subnetwork = "${var.subnet}"

    # access_config {}
  }

  metadata {
    startup-script = "${file("${path.module}/init.sh")}"
  }
}


resource "null_resource" "wait_for_startup" {
  provisioner "local-exec" {
    command = "${path.module}/wait-for-startup.sh ${var.name} ${google_compute_instance.redis_vm.network_interface.0.network_ip}"
  }

  depends_on = ["google_compute_instance.redis_vm"]
}