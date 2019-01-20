#! /bin/bash
docker run \
    -v $(pwd):/var/loadtest \
    -v $SSH_AUTH_SOCK:/ssh-agent -e SSH_AUTH_SOCK=/ssh-agent \
    -v /home/mlychagin/dev/python/hlcupdocs/data/test_accounts_241218/ammo/phase_2_post.ammo:/var/loadtest/phase_2_post.ammo \
    --net host \
    -it direvius/yandex-tank
