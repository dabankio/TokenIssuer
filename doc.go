// package TokenIssuer -- Issue Your ERC-20 Token in a Minute
package TokenIssuer

//go:generate solc -o contract --abi --bin --optimize --overwrite contract/token.sol
//go:generate abigen --abi contract/Token.abi --bin contract/Token.bin --pkg TokenIssuer --type Token --out token.go
