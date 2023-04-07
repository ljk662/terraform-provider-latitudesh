---
page_title: "latitudesh_server Resource - terraform-provider-latitudesh"
subcategory: ""
description: |-
  
---

# latitudesh_server (Resource)



## Example usage

```terraform
resource "latitudesh_server" "server" {
  hostname         = "terraform.latitude.sh"
  operating_system = "ubuntu_22_04_x64_lts"
  plan             = data.latitudesh_plan.plan.slug
  project          = latitudesh_project.project.id      # You can use the project id or slug
  site             = data.latitudesh_region.region.slug # You can use the site id or slug
  ssh_keys         = [latitudesh_ssh_key.ssh_key.id]
  user_data        = latitudesh_user_data.user_data.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hostname` (String) The server hostname
- `operating_system` (String) The server OS
- `plan` (String) The server plan
- `project` (String) The id or slug of the project
- `site` (String) The server site

### Optional

- `ssh_keys` (List of Number) List of server SSH key ids
- `user_data` (Number) The id of user data to set on the server

### Read-Only

- `created` (String) The timestamp for when the server was created
- `id` (String) The ID of this resource.
- `primary_ipv4` (String) The server IP address
- `updated` (String) The timestamp for the last time the server was updated

## Import
Server can be imported using the serverID, e.g.,

```sh
$ terraform import latitudesh_server.server serverID
```