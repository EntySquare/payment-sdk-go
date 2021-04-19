# payment-sdk-go
For golang usage only to custom payment usage

## we use tls/ssl to encrypt payment transaction from given CA


### from serverside we use keys generated based on elliptic curve like
    
    openssl ecparam -genkey -name secp384r1 -out foo.key 

### we deploy CA use underlying to generate a cert for ten years
    
    openssl req -new -x509 -sha256 -key foo.key -out foo.pem -days 3650

