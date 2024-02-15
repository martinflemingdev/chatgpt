curl -k -X POST -H "Content-Type: application/json" -d '{"key1": "value1", "key2": {"subkey1": "subvalue1", "subkey2": "subvalue2"}}' https://your-website.com/endpoint

# With auth
curl -k -X POST -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_TOKEN_HERE" -d '{"key1": "value1", "key2": {"subkey1": "subvalue1", "subkey2": "subvalue2"}}' https://your-website.com/endpoint
