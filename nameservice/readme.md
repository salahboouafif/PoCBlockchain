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

 the "startport serve" command  is detailed in the file init.sh

## Test the application 
 First check the accounts to ensure they have funds 
 
    nameservicecli query account $(nameservicecli keys show user1 -a)
    nameservicecli query account $(nameservicecli keys show user2 -a)

 Buy your first name using your coins from the genesis file

    nameservicecli tx nameservice buy-name user1.id 5nametoken --from user1

 Set the value for the name you just bought

    nameservicecli tx nameservice set-name user1.id 8.8.8.8 --from user1
    
 User2 buys name from user1

    nameservicecli tx nameservice buy-name user1.id 10nametoken --from alice   
    
 To delete the name, user 2 executes 
 
    nameservicecli tx nameservice delete-name user1.id --from user2
  ## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
