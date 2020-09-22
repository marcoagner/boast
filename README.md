# BOAST
_The BOAST Outpost for AppSec Testing (v0.1.0)_

BOAST is a server built to receive and report Out-of-Band Application Security Testing reactions.

<p align="center">
  <img src="./docs/boast.png" alt="BOAST overview">
</p>

Some application security tests will only cause out-of-band reactions from the tested
applications. This means that these reactions will not be sent as a response to the
testing client and, due to their nature, they will remain unseen when the client is
behind a third-party NAT. For the purpose of being able to clearly see those reactions,
another piece is needed. A piece that, not limited by a third-party NAT, is freely
reachable on the Internet and can also speak the received protocols in multiple ports
for maximum impact. BOAST is this piece.

BOAST features DNS, HTTP, and HTTPS protocol receivers with support for multiple
simultaneous ports for each receiver. And implementing protocol receivers for new
protocols or to better suit your needs is almost as simple as implementing the protocol
interaction itself.

## Documentation

The project is documented [here](https://github.com/marcoagner/boast/tree/master/docs).
