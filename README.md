# Terraform TimeCrowd Provider

## Install

```
# macos
$ make install
```

## Usage

```
terraform {
  required_providers {
    timecrowd = {
      version = "0.1"
      source  = "hashicorp.com/machamp/timecrowd"
    }
  }
}

data "timecrowd_team" "team" {
  id = 1
}

data "timecrowd_category" "category" {
  team_id = data.timecrowd_team.team.id
  id      = 1
}

resource "timecrowd_category" "category" {
  team_id   = data.timecrowd_team.team.id
  parent_id = data.timecrowd_category.category.id
  title     = "category title"
  color     = 5
}
```

```
terraform init && terraform plan
terraform apply
```