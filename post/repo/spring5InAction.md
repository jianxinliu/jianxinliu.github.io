


# spring in action

## what is spring

核心是一个**容器**，也被称为上下文（context），用于**创建和管理组件**（bean，component），spring 负责将组件**装配**起来，形成一个完成的应用。

装配使用依赖注入（DI或控制反转（IoC））实现。程序员不用过多干预组件如何耦合，只需要设定耦合关系， spring 便会在合适的时机以合适的方式将需要的组件进行装配。

```java
@RestController
public class UserController {
	// 声明一个依赖，而不需要手动创建和管理该依赖的生命周期
	// 通过属性设置方法注入
	// 需要将 UserService 的实现类声明为组件（@Component\@Service）
	@Autowired
	private UserService userService;
	
	// 或者通过构造器注入，二者选其一即可。通常建议使用构造器注入
	@Autowired
	public UserController(UserService userService){
		this.userService = userService;
	}
}
```

## View Controller

视图控制器，更简单的声明仅仅做视图转发的控制器

```java
@Configuration
class WebConfig implements WebMvcConfigurer {
	@Override
	public void addViewControllers(ViewControllerRegistry registry){
		registry.addViewController("/").setViewName("home");
		// ...
	}
}
```

## bean validation

声明式校验
```java
class Dto{
	@NotNull
	@NotBlank(message="")
	@Size(min=1,message="not less then 1")
	@Pattern(regexp="")
	@Digits(integers=3,fraction=0,message="")
}
(@Valid Dto dto,Errors errs){
	if(errs.hasErrors()){
		// valid failed
		// errors.getAllErrors().forEach(err -> sj.add(err.getDefaultMessage()))
	}
}
```

## spring data

### 配置数据源



### jdbc -> jdbcTemplate

```java
private JdbsTemplate jdbc;

Student find(String id){
	// jdbc.query(sql,orm)
	// jdbc.update(sql,...params)
	return jdbc.queryForObject("select * from student where id=?",this::orm,id)
}

Student orm(ResultSet rs,int rowNum) throw SQLException{
	return new Student(rs.getString("name"),...)
}
```

sping boot 自动预定义与预加载数据

在根路径下（src/main/resources）放置 schema.sql 和 data.sql 即可在项目启动时自动执行

### jpa

通过在 Repostiry 接口中添加特定命名的方法自动增加实现（DSL）
如添加 `findByxxx` 方法，即可自动按照 Domain 中的 xxx 字段查找

还可在方法上添加 `@Query("order by xxx where xxx")` 注解，在其中添加更复杂的逻辑

## 安全、权限

spring security（太有限，太刻板，不够灵活，往往都是自己写权限控制逻辑）

权限控制

## 自动配置

`@Bean` 自动装配

Spring的环境抽象是各种配置属性的一站式服务。它抽取了原始的属性，这样需要这些属性的bean就可以从Spring本身中获取了。
Spring环境会拉取多个属性源，包括：
•JVM系统属性；
•操作系统环境变量； export SERVER_PORT=9090
•命令行参数；java -jar xxx --server.port=9090
•应用属性配置文件。 application.properties or application.yml
它会将这些属性聚合到一个源中，通过这个源可以注入到Spring的bean中。

## 消费 REST 服务

RestTemplate 同步 Rest 客户端
WebClient 异步 Rest 客户端

```java
RestTemplate rest = new RestTemplate();

// 自动将结果解析成对象，参数可以使用可变参数依次指定，也可以使用 Map 指定具名参数。
rest.getForObject(url,<Object>.class,...params);

rest.getForEntity() // 类似 getForObject,只是返回的是一个包含了响应体的更丰富的对象 ResponseEntity,从中可以获得响应头等信息

rest.put()
rest.delete()
rest.postForObject()
```

## exception handler

## log

lombok:@Slf4j => log.*

advisor