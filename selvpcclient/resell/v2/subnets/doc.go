/*
Package subnets provides the ability to retrieve and manage subnets through
the Resell v2 API.

Example of getting a single subnet referenced by its id

  subnet, _, err := subnets.Get(context, resellClient, subnetID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(subnet)

Example of getting all subnets

  allSubnets, _, err := subnets.List(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, subnet := range allSubnet {
    fmt.Println(subnet)
  }
*/
package subnets