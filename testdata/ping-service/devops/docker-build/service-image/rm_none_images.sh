# /bin/sh

export NONE_IMAGES=$(docker images --filter "dangling=true" -q)
if [ "${NONE_IMAGES}" != "" ]; then
		docker image rm ${NONE_IMAGES} || true
fi