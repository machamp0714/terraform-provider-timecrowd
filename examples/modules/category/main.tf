terraform {
  required_providers {
    timecrowd = {
      version = "0.1"
      source  = "hashicorp.com/edu/timecrowd"
    }
  }
}

resource "timecrowd_category" "this" {
  team_id = var.team_id

  title     = var.title
  color     = var.color
  parent_id = var.parent_id
  position  = var.position
}
