    ctr=0
    max_retries=10
    while [ $(curl -I -u "root:123qwe123" ocalhost:10080/users/sign_in | head -1 | grep 200 | wc -l | tr -d ' ') -ne 1 ]; do
      ctr=`expr $ctr + 1`
      echo waiting for gitlab to come up
      sleep 5
      if [ $ctr == $max_retries ]; then
        echo "max retries (${max_retries}) reached"
        exit 1
      fi
    done
