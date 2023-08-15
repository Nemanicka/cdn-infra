$TTL 1d
$ORIGIN perfecthw.com.ua.

@ IN  SOA   ns1 root (
      1     ;Serial
      1d    ;Refresh
      1h    ;Retry
      1w    ;Expire
      1h    ;Minimum
  )

@      IN     A       127.0.0.1

@      IN     NS      ns1.perfecthw.com.ua.
ns1    IN     A       127.0.0.10

@       IN    NS      ns2.perfecthw.com.ua.
ns2     IN    A       127.0.0.11

www     IN    CNAME   perfecthw.com.ua.
