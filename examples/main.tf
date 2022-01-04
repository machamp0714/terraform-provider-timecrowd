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

data "timecrowd_category" "terraform" {
  team_id = data.timecrowd_team.team.id
  id      = 1
}

module "aws_category" {
  source = "./modules/category"

  team_id   = data.timecrowd_team.team.id
  title     = "aws"
  parent_id = data.timecrowd_category.terraform.id
  position  = 1
}
