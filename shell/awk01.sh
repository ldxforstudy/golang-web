# 基于log统计API访问次数
#
#

{
    if($2 ~/\_com\_request\_out/) {
        split($2, A, /\|\|/);
        map[A["9"]]++
    }
}

END {
    for(i in map){
        print map[i], "\t", i
    }
}