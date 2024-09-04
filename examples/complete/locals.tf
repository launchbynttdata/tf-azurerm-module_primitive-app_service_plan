locals {
  resource_group_name = module.resource_names["resource_group"].minimal_random_suffix
  service_plan_name   = module.resource_names["app_service_plan"].minimal_random_suffix
}
