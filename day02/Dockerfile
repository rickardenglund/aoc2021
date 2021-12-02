FROM codesimple/elm:0.19
WORKDIR /src
COPY ./ .
RUN elm make src/Day.elm --output /dev/null

WORKDIR /src/src
ENTRYPOINT echo "$PART" >> /src/prog && cat /src/prog | elm repl
