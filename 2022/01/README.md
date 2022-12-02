### A

#This caps out at 252, but there are more...

```
awk -v RS= '{print > ("" NR ".txt")}' input
```

#find contents of 252.txt, and get the rest into input2

```
awk -v RS= '{print > ("aaaa_" NR ".txt")}' input2
```

#rename last batch

```
for i in $(ls aaaa_*.txt); do
    x=$(echo $i | sed 's/\.txt//' | sed 's/aaaa_//')
    y=$((252 + $x))
    #echo $y.txt
    mv $i $y.txt
done
```

#calculate totals

for file in $(ls *.txt); do
    awk -v file="$file" '{ sum += $1 } END { print sum " " file}' $file
done > output

# get largest calories
sort -rn output | head -1 | cut -d' ' -f1

### B

# get top 3 calories

sort -rn output | head -3 | cut -d' ' -f1 > input3

# total top 3
paste -s -d+ input3 | bc

