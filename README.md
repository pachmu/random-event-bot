## for docker
1. build image

     docker build -t bot -f Dockerfile .

2. create a config.yml file at ./
  
  token: some_token
  
3. run

     docker run --rm -dit bot