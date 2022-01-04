terraform {
  required_providers {
    timecrowd = {
      version = "0.1"
      source  = "hashicorp.com/edu/timecrowd"
    }
  }
}

data "timecrowd_team" "team" {
  id = 21648
}

data "timecrowd_category" "category" {
  team_id = data.timecrowd_team.team.id
  id      = 44163
}

output "team" {
  value = data.timecrowd_team.team
}

output "category" {
  value = data.timecrowd_category.category
}
