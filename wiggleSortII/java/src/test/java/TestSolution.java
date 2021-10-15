import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.lang.reflect.Array;
import java.util.stream.Stream;

import static org.junit.jupiter.params.provider.Arguments.arguments;

public class TestSolution {

    Solution solution = new Solution();

    @ParameterizedTest
    @MethodSource("kthnums")
    public void TestFindKth(int[] nums, int k, int expect){

        var actual = Solution.findKth(nums, 0, nums.length-1, k);
        var a2 = Solution.findKthLargest(nums, k);
        Assertions.assertEquals(expect, actual);
    }
    private static Stream<Arguments> kthnums() {
        return Stream.of(
                arguments(new int[]{5,3,1,2,6,7,8,5,5}, 4, 1)
        );
    }

}
