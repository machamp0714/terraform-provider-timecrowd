terraform {
  required_providers {
    timecrowd = {
      version = "0.1"
      source  = "hashicorp.com/edu/timecrowd"
    }
  }
}

data "timecrowd_team" "team" {
  id = 23260
}

data "timecrowd_category" "category" {
  team_id = data.timecrowd_team.team.id
  id      = 1713677
}

output "team" {
  value = data.timecrowd_team.team
}

output "category" {
  value = data.timecrowd_category.category
}

resource "timecrowd_category" "example" {
  team_id = data.timecrowd_team.team.id

  title     = "aws"
  color     = 3
  parent_id = data.timecrowd_category.category.id
  position  = 0
}

output "aws_category" {
  value = timecrowd_category.example
}
