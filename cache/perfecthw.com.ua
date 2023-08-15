$TTL 1d
$ORIGIN perfecthw.com.ua.

@ IN  SOA   ns1 root (
      1     ;Serial
      1d    ;Refresh
      1h    ;Retry
      1w    ;Expire
      1h    ;Minimum
  )

IN NS lb1.perfecthw.com.ua
lb1 IN A 127.0.0.10 

IN NS lb2.perfecthw.com.ua
lb1 IN A 127.0.0.11
