#!/bin/sh
echo "window._env_ = { \"REACT_APP_API_URL\":\""$REACT_APP_API_URL"\", \"REACT_APP_WS_URL\":\""$REACT_APP_WS_URL"\" }" > ./public/env-config.js