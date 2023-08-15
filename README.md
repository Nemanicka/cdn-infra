# Setup

Overall I have 3 different load balancers:
1. nginx-lb-1 using least_conn directive
2. nginx-lb-2 using least_time directive, it an image recompiled with a module that enables response time comparison. The image is pushed to docker hub.
3. lb-3 - a primitive golang reverse proxy that also counts number of hops to the servers and picks tha best one. In our setup, the number of hops is always 1, so the first server is always returned

In docker compose either nginx-lb-2 or lb-3 should be commented since they both linked to node-3 and node 4.

Overall, the setup can be checked via 
curl  http://perfecthw.com.ua/dog.jpg

Unfortunatelly, bind9 somewhy doesn't recognize my Country.mmdb, therefore, I didn't find a way to exmulate anycast, although the rest of the config works pretty ok and the only difference in my setup that all requests will be redirected to lb-1
