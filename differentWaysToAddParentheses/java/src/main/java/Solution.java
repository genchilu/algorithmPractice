import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> diffWaysToCompute(String expression) {
        var res = new ArrayList<Integer>();
        for (int i =0;i<expression.length();i++) {
            if(expression.charAt(i) == '+' || expression.charAt(i) == '-' || expression.charAt(i) == '*') {
                var p1 = expression.substring(0,i);
                var res1 = diffWaysToCompute(p1);

                var p2 = expression.substring(i+1);
                var res2 = diffWaysToCompute(p2);

                for (var a:res1) {
                    for (var b:res2) {
                        switch (expression.charAt(i)) {
                            case '+':
                                res.add(a+b);
                                break;
                            case '-':
                                res.add(a-b);
                                break;
                            case '*':
                                res.add(a*b);
                                break;
                        }
                    }
                }

            }
        }

        if (res.size()==0) {
            res.add(Integer.parseInt(expression));
        }

        return res;
    }
}

