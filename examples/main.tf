terraform {
  required_providers {
    timecrowd = {
      version = "0.1"
      source = "hashicorp.com/edu/timecrowd"
    }
  }
}

data "timecrowd_team" "team" {
  id = 21648
}

output "ruffnote_team" {
  value = data.timecrowd_team.team
}
