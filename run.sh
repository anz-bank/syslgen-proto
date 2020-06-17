shopt -s nullglob
echo "++++++++++++++++++++++++++++++++++++++++++++"
echo "       Proto generation from Sysl           " 
echo "++++++++++++++++++++++++++++++++++++++++++++"
for file in demo/*.sysl; do
    echo "Reading Sysl file \"$file\""
    echo "Generating json for \"$file\""
    read -p "Enter the target sysl appname, comma separate if more than one : " app
    
    IFS=',' read -ra apps <<< "$app"
    for i in "${apps[@]}"
    do
        read -p "Enter the target sysl package for "$i": " package
        rm -rf "demo/$i.json"
        rm -rf "gen/proto/$package"
        mkdir "gen/proto/$package"
        sysl pb --mode json $file >> "demo/$i.json"
    
        echo "Now generating proto files for $i"
        echo "Running arrai.ai to generate proto file"
        arrai r transform/grpcgen.arrai $i "demo/$i.json" >> "gen/proto/$package/$i.proto"
        protoc -I=gen/proto/$package --go_out=gen/proto/$package "gen/proto/$package/$i.proto"
    done
done
echo "++++++++++++++++++++++++++++++++++++++++++++"
echo "                    End                     " 
echo "++++++++++++++++++++++++++++++++++++++++++++"
