import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class DetectSquaresTest {

    @Test
    public void TestDev1() {
        var target = new DetectSquares();

        target.add(new int[]{3,10});
        target.add(new int[]{11,2});
        target.add(new int[]{3,2});

        var actual = target.count(new int[]{11,10});
        Assertions.assertEquals(1, actual);

        actual = target.count(new int[]{14,8});
        Assertions.assertEquals(0, actual);

        target.add(new int[]{11,2});
        actual = target.count(new int[]{11,10});
        Assertions.assertEquals(2, actual);

    }
}
