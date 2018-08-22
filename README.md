# golang simple rest api bank



# testing REST API

# POST http://localhost:7000/account/create
id=1
name=Mark
amount=8.99

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:50:13 GMT
Content-Length: 0

<Response body is empty>

================================================

# POST http://localhost:7000/account/create
id=2
name=Mark2
amount=10.00

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:50:34 GMT
Content-Length: 0

<Response body is empty>

================================================

# GET http://localhost:7000/account/get/1

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:51:37 GMT
Content-Length: 49
Content-Type: text/plain; charset=utf-8

{"Id":1,"Name":"Mark","Amount":{"Value":"8.99"}}

===============================================

# GET http://localhost:7000/account/get/2

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:52:53 GMT
Content-Length: 48
Content-Type: text/plain; charset=utf-8

{"Id":2,"Name":"Mark2","Amount":{"Value":"10"}}

===============================================

# GET http://localhost:7000/account/get/1/amount

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:57:46 GMT
Content-Length: 7
Content-Type: text/plain; charset=utf-8

"8.99"

===============================================

# GET http://localhost:7000/account/get/2/amount

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 00:58:47 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

"10"

===============================================

# PUT http://localhost:7000/account/pay
sendAccountId=1
receiveAccountId=2
summ=5.47

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 01:00:55 GMT
Content-Length: 0

<Response body is empty>

==============================================

# GET http://localhost:7000/account/get/1

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 01:02:08 GMT
Content-Length: 49
Content-Type: text/plain; charset=utf-8

{"Id":1,"Name":"Mark","Amount":{"Value":"3.52"}}

===============================================

# GET http://localhost:7000/account/get/2

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 01:03:00 GMT
Content-Length: 51
Content-Type: text/plain; charset=utf-8

{"Id":2,"Name":"Mark2","Amount":{"Value":"15.47"}}

==============================================

# GET http://localhost:7000/account/get/1/amount

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 01:04:06 GMT
Content-Length: 7
Content-Type: text/plain; charset=utf-8

"3.52"

==============================================

# GET http://localhost:7000/account/get/2/amount

HTTP/1.1 200 OK
Date: Wed, 22 Aug 2018 01:04:25 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8

"15.47"
