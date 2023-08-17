# Setup

Overall I have 3 different load balancers:
1. nginx-lb-1 using least_conn directive
2. nginx-lb-2 using least_time directive, it is an image recompiled with a module that enables response time comparison. The image is pushed to the docker hub.
3. lb-3 - a primitive Golang reverse proxy that also counts the number of hops to the servers and picks the best one. In our setup, the number of hops is always 1, so the first server is always returned

In docker compose either nginx-lb-2 or lb-3 should be commented on since they are both linked to node-3 and node 4.

Overall, the setup can be checked via 
curl  http://perfecthw.com.ua/dog.jpg

Unfortunately, bind9 somewhy doesn't recognize my Country.mmdb, therefore, I didn't find a way to emulate anycast, although the rest of the config works pretty ok and the only difference in my setup is that all requests will be redirected to lb-1

# Conclusions

It's pretty hard to make any conclusions in this setup since there's very little apple-to-apple comparison, and hop-based balancing is completely artificial.
Both least_time and least_connection are working pretty ok, but on the local machine this might not be representative tests.
