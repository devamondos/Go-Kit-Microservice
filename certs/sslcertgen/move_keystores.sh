if [ ! -d ../ca ]; then
    mkdir ../ca
else
    rm -f ../ca/*
fi

if [ ! -d ../client ]; then
    mkdir ../client
else
    rm -f ../client/*
fi

if [ ! -d ../server ]; then
    mkdir ../server
else
    rm -f ../server/*
fi

mv cacert.jks ../ca
mv cacert.pem ../ca
mv cakey.pem ../ca

mv 100001.pem ../server
mv server.jks ../server
mv servercert.pem ../server
mv serverkey.pem ../server
mv serverreq.pem ../server

mv 100002.pem ../client
mv client.jks ../client
mv clientcert.pem ../client
mv clientcert.p12 ../client
mv clientkey.pem ../client
mv clientreq.pem ../client

