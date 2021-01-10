# nameservice

**nameservice** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

## Build the application

    make install
## run the application
```
starport serve
```
`serve` command installs dependencies, initializes and runs the application.

Instead of  `starport serve` we can use : 
 
      ./init.sh
      nameserviced start

## Test the application 
 First check the accounts to ensure they have funds 
 
    nameservicecli query account $(nameservicecli keys show user1 -a)
    nameservicecli query account $(nameservicecli keys show user2 -a)

 Buy your first name using your coins from the genesis file

    nameservicecli tx nameservice buy-name name_test 5nametoken --from user1

 Set the value for the name you just bought

    nameservicecli tx nameservice set-name name_test 8.8.8.8 --from user1
    
 User2 buys name from user1

    nameservicecli tx nameservice buy-name name_test 20nametoken --from user2   
    
 To delete the name, user 2 executes 
 
    nameservicecli tx nameservice delete-name name_test --from user2
  ## Configure

The initial configuration of client accounts is specified in `config.yml`.


## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
