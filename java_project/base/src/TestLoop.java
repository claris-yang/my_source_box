public class TestLoop {
    public void testLoop() {
        LABLE1:
        for(int i = 0 ; i < 10; i++) {

            for(int j =0; j < 100; j++) {
                if(j > 50) continue LABLE1;
                System.out.println(j);
            }
        }
    }
}
