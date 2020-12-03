# td

Telegram client implementation in go.

* [Security policy](.github/SECURITY.md)
* [Support](.github/SUPPORT.md)

## Status

Work in progress. Only go1.15 is supported and no backward compatibility is provided.

Goal of this project is to implement Telegram client while
providing building blocks for the other client or even server
implementation without performance bottlenecks.

This project is fully non-commercial and not affiliated with any commercial organization
(including Telegram LLC).

## Features

* Full MTProto 2.0 implementation in go
* Code for Telegram types automatically generated by `./cmd/gotdgen` (based on [gotd/tl](https://github.com/gotd/tl) parser)
* Pluggable session storage
* Vendored Telegram public keys that are kept up-to-date
* Testing and fuzzing applied to improve durability
* No runtime reflection overhead
* Secure PRNG used for crypto

## Example

You can see `cmd/gotdecho` for echo bot example.

### Calling MTProto directly

You can use generated `tg.Client` that allows calling any MTProto methods
directly.

```go
client, err := telegram.Dial(ctx, telegram.Options{
    Addr:   "149.154.167.40:443",

    // Grab these from https://my.telegram.org/apps.
    // Never share it or hardcode!
    AppID:   appID,
    AppHash: appHash,
})
if err != nil {
    panic(err)
}
// Grab token from @BotFather.
if err := client.BotLogin(ctx, "token:12345"); err != nil {
    panic(err)
}
// updates.getState#edd4882a
state, err := tg.NewClient(client).UpdatesGetState(ctx, &tg.UpdatesGetStateRequest{})
if err != nil {
    panic(err)
}
// Got state: &{Pts:197 Qts:0 Date:1606855030 Seq:1 UnreadCount:106}
```

### Generated code

Code output of `gotdgen` contains references to TL types, examples and URL to
official documentation.

For example, the [auth.Authorization](https://core.telegram.org/type/auth.Authorization) type in `tg/tl_auth_authorization_gen.go`:

```go
// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthAuthorization: // auth.authorization#cd050916
//  case *AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthAuthorizationClass
}
```
Also, the corresponding [auth.signIn](https://core.telegram.org/method/auth.signIn) method:
```go
// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {}
```

## Reference

The MTProto protocol description is [hosted](https://core.telegram.org/mtproto#general-description) by Telegram.

Most important parts for client impelemtations:
* [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for client software developers

Current implementation does not conform to security guidelines and should be used only
as reference or for testing.

## Prior art

* [Lonami/grammers](https://github.com/Lonami/grammers) (Great Telegram client in Rust, many test vectors were used as reference)
* [sdidyk/mtproto](https://github.com/sdidyk/mtproto), [cjongseok/mtproto](https://github.com/cjongseok/mtproto), [xelaj/mtproto](https://github.com/xelaj/mtproto)  (MTProto 1.0 in go)

## License
MIT License

Created by Aleksandr (ernado) Razumov

2020
