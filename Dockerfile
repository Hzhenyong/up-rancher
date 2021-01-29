FROM neilpang/acme.sh:latest

COPY ./main /root
RUN echo "0 1 1 */2 * /root/main 2 >& 1 >> /dev/null &"


ENTRYPOINT ["/entry.sh"]
CMD ["--help"]