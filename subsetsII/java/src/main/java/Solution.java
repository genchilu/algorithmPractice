import java.util.*;
import java.util.stream.Collectors;

class Solution {
    public List<List<Integer>> subsetsWithDup(int[] nums) {

        var res = new ArrayList<List<Integer>>();
        Map<Integer, Integer> map = new LinkedHashMap<>();
        for (var i:nums) {
            var count = map.getOrDefault(i, 0);
            map.put(i, count+1);
        }
        Integer[] keysArray = map.keySet().toArray(new Integer[map.size()]);

        for(int i = 0;i<=nums.length;i++){
            pick(map, keysArray, 0,i, new ArrayList<>(), res);
        }
        return res;
    }

    public void pick(Map<Integer, Integer> map, Integer[] keysArray, int idx, int l, List<Integer> curPick, List<List<Integer>> result) {
        if (l==curPick.size()) {
            result.add(curPick);
        } else {
            for(int i = idx;i<keysArray.length;i++){

                var val = keysArray[i];
                for(int j = 1;j<=map.get(val);j++){

                    var cloneCurPick = curPick.stream().collect(Collectors.toList());
                    for(int k=1;k<=j;k++) {
                        cloneCurPick.add(val);
                    }

                    if(cloneCurPick.size()>l) {
                        break;
                    }

//                    System.out.println("==== l: " + l + "  ==== i: " + i + " ==== j: " + j + "==== idx : " + idx);
//                    for(var s:cloneCurPick) {
//                        System.out.print(s + ",");
//                    }
//                    System.out.println("");

                    pick(map, keysArray, i+1, l, cloneCurPick,result);
                }
            }
        }
    }
}
