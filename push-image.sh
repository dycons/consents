#!/bin/bash

################################################################################
# Help                                                                         #
################################################################################
help ()
{
   # Display Help
   echo "Build, tag, and push a docker image."
   echo
   echo "Usage:"
   echo "   ./push-image.sh [options] IMAGE PATCH"
   echo "   docker login && ./push-image.sh [options] IMAGE PATCH"
   echo "It may be necessary to run the following command before running this script:"
   echo "   docker login"
   echo "Mandatory positional arguments:"
   echo "   IMAGE   The name of the image, NOT including the username"
   echo "   PATCH   patch version (semantic)"
   echo "Options:"
   echo "   -h      Display this help text"
   echo "   -b      Build image prior to push"
   echo "   -M      Major version (semantic). Default: 1"
   echo "   -m      minor version (semantic). Default: 1"
   echo "   -f      Dockerfile filename. Default: ./Dockerfile"
   echo "   -u      Docker username. Default: katpavlov"
   echo
}

################################################################################
################################################################################
# Main program                                                                 #
################################################################################
################################################################################

# Default Docker username. Can overwrite with -u argument.
username="katpavlov"
# Default image version. Can overwrite with -M -m -p arguments.
Major="0"
minor="0"
# Default Dockerfile name. Can overwrite with -f argument.
Dockerfile="./Dockerfile"
# Flag for whether to build the image prior to attempting push
build=false

# Optionally overwrite image version, Dockerfile, or Docker username
while getopts ":hbM:m:f:u:" opt; do
  case $opt in
    h)  help
        exit
        ;;
    b)  build=true
        ;;
    M)  Major="$OPTARG"
        ;;
    m)  minor="$OPTARG"
        ;;
    f)  Dockerfile="$OPTARG"
        ;;
    u)  username="$OPTARG"
        ;;
    \?) echo "Invalid option -$OPTARG" >&2
        ;;
  esac
done
shift $((OPTIND -1))

# Image name (mandatory)
image="$1"
# Image patch version (mandatory).
# This should be a single number; major and minor can be defaulted or defined in the options.
patch="$2"


# Optional build-image step
if [ "$build" = true ]
then
  docker build -t $username/$image -f $Dockerfile .
fi

# Push image tagged with the semantic versioning system described here: 
#       https://medium.com/@mccode/using-semantic-versioning-for-docker-image-tags-dfde8be06699
docker tag $username/$image $username/$image:latest
docker push $username/$image:latest

docker tag $username/$image:latest $username/$image:$Major
docker push $username/$image:$Major

docker tag $username/$image:latest $username/$image:$Major.$minor
docker push $username/$image:$Major.$minor

docker tag $username/$image:latest $username/$image:$Major.$minor.$patch
docker push $username/$image:$Major.$minor.$patch
