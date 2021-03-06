// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


// Run:
//  terraform apply [-var agones_version="1.2.0"]
// to install "1.2.0" version of agones

variable "agones_version" {
  default = "1.2.0"
}

variable "cluster_name" {
  default = "agones-cluster"
}

variable "region" {
  default = "us-west-2"
}

variable "node_count" {
  default = "4"
}

provider "aws" {
  version = "~> 2.8"
  region  = var.region
}

variable "machine_type" { default = "t2.large" }

module "eks_cluster" {
  source = "git::https://github.com/googleforgames/agones.git//install/terraform/modules/eks/?ref=master"

  machine_type = "${var.machine_type}"
  cluster_name = "${var.cluster_name}"
  node_count   = "${var.node_count}"
  region       = "${var.region}"
}

data "aws_eks_cluster_auth" "example" {
  name = "${var.cluster_name}"
}

// Next Helm module cause "terraform destroy" timeout, unless helm release would be deleted first.
// Therefore "helm delete --purge agones" should be executed from the CLI before executing "terraform destroy".
module "helm_agones" {
  source = "git::https://github.com/googleforgames/agones.git//install/terraform/modules/helm/?ref=master"

  udp_expose             = "false"
  agones_version         = "${var.agones_version}"
  values_file            = ""
  chart                  = "agones"
  host                   = "${module.eks_cluster.host}"
  token                  = "${data.aws_eks_cluster_auth.example.token}"
  cluster_ca_certificate = "${module.eks_cluster.cluster_ca_certificate}"
}

output "host" {
  value = "${module.eks_cluster.host}"
}
output "cluster_ca_certificate" {
  value = "${module.eks_cluster.cluster_ca_certificate}"
}
