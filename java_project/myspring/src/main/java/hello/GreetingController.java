package hello;

import java.util.HashMap;
import java.util.Map;
import java.io.Reader;
import java.util.concurrent.atomic.AtomicLong;

import net.sf.json.JSONObject;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import redis.clients.jedis.Jedis;

@RestController
public class GreetingController {

    private static final String template = "Hello, %s!";
    private final AtomicLong counter = new AtomicLong();

    @RequestMapping("/greeting")
    public Greeting greeting(@RequestParam(value="name", defaultValue="World") String name) {

        Jedis jedis = new Jedis("127.0.0.1");
        jedis.set("foo", name);
        String value = jedis.get("foo");

        return new Greeting(counter.incrementAndGet(),
                String.format(template, name));

    }
}
