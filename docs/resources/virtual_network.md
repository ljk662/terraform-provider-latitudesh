---
page_title: "latitudesh_virtual_network Resource - terraform-provider-latitudesh"
subcategory: ""
description: |-
  
---

# latitudesh_virtual_network (Resource)



## Example usage

```terraform
resource "latitudesh_virtual_network" "virtual_network" {
  description      = "Virtual Network description"
  site             = data.latitudesh_region.region.slug # You can use the site id or slug
  project          = latitudesh_project.project.id      # You can use the project id or slug
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The virtual network description
- `project` (String) The project id or slug
- `site` (String) The site id or slug

### Read-Only

- `assignments_count` (Number) Amount of devices assigned to the virtual network
- `id` (String) The ID of this resource.
- `vid` (Number) The vlan ID of the virtual network

## Import
VirtualNetwork can be imported using the VirtualNetworkID along with the projectID that contains the VirtualNetwork, e.g.,

```sh
$ terraform import latitudesh_virtual_network.virtual_network projectID:VirtualNetworkID
```