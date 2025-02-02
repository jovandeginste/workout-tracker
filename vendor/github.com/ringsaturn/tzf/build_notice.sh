echo "NOTICE" > NOTICE
echo "" >> NOTICE
echo "This product includes software developed by [github.com/ringsaturn/tzf]." >> NOTICE
echo "------------------------------------------------------------------------" >> NOTICE

for license_file in $(find THIRD_PARTY_LICENSES -name LICENSE); do
    dependency=$(basename $(dirname $license_file))
    license_type=$(grep -E -o "(MIT|BSD|Apache)" $license_file | head -1)
    copyright=$(grep -i "copyright" $license_file | head -1)
    
    echo "" >> NOTICE
    echo "$dependency" >> NOTICE
    echo "License: $license_type" >> NOTICE
    echo "$copyright" >> NOTICE
done
