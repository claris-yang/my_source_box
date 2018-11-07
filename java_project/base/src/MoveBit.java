public class MoveBit {
    public void testMoveBit(int i, int moveBitNums) {
        System.out.println(i >>>= moveBitNums);
    }

    public void testMoveLongBit(long i , int moveBitNums) {
        System.out.println(i >>>= moveBitNums);
    }
}
