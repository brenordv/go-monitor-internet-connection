#!/bin/bash
sudo docker run -d --restart=unless-stopped --name conn_monitor -it raccoon_conn_monitor2:latest