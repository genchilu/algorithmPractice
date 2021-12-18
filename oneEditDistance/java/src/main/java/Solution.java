class Solution {
    public boolean isOneEditDistance(String s, String t) {

        if(s==null || t==null) {
            return false;
        }

        if(s.equals(t)) {
            return false;
        }

        if(s.length() > t.length()) {
            return checkDelete(s, t);
        }

        if(t.length()>s.length()) {
            return checkDelete(t,s);
        }

        return checkReplace(s,t);
    }

    public boolean checkDelete(String longer, String shoter) {
        if(longer.length() - shoter.length() !=1) {
            return false;
        }

        for(int i =0;i<shoter.length();i++){
            if(longer.charAt(i)!=shoter.charAt(i)) {
                return longer.substring(i+1).equals(shoter.substring(i));
            }
        }

        return true;
    }

    public boolean checkReplace(String s, String t) {
        if(s.length()!=t.length()) {
            return false;
        }

        for(int i =0;i<s.length();i++){
            if(s.charAt(i) != t.charAt(i)) {
                return s.substring(i+1).equals(t.substring(i+1));
            }
        }

        return true;
    }
}