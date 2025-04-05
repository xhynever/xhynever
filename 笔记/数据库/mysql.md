## 单表查询

1）请编写sql语句对年龄进行升序排列

select * from afinfo order by birth;

2）请编写sql语句查询对“徐”姓开头的人员名单

select * from afinfo where name like '徐%';

3）请编写sql语句修改“陈晓”的年龄为“45”

update afinfo set age=45 and birth=birth-YEAR(45) where name="陈晓";

4）请编写sql删除王芳芳这表数据记录。

delete from afinfo where name="王芳芳";


## 学生表和成绩表查询

1）查询出所有学生信息，SQL怎么编写？

select * from stu;


2）新学生小明，学号为005，需要将信息写入学生信息表，SQL语句怎么编写？

insert into stu values ("小明",005);
 

3）李四语文成绩被登记错误，成绩实际为85分，更新到考试信息表中，SQL语句怎么编写？

update exam set score=85 where id=(select id from stu 
where name="李四") and subject="语文";


4）查询出各科成绩的平均成绩，显示字段为：学科、平均分，SQL怎么编写？

select subject,avg(score) from exam group by subject;


5）查询出所有学生各科成绩，显示字段为：姓名、学号、学科、成绩，并以学号与学科排序，没有成绩的学生也需要列出，SQL怎么编写？

select s.name,s.id,e.subject,e.score from stu s 
left join exam e 
on s.id=e.id 
order by id,subject;


6）查询出单科成绩最高的，显示字段为：姓名、学号、学科、成绩，SQL怎么编写？

select s.name,s.id,e.subject,e.score from stu s 
join exam e 
on s.id=e.id 
where (e.subject,e.score) 
in (select subject,max(score) from exam group by subject);


7）列出每位学生的各科成绩，要求输出格式：姓名、学号、语文成绩、数学成绩、英语成绩，SQL怎么编写？

## 根据要求写sql
1、查询“001”课程比“002”课程成绩高的所有学生的学号。

select a.s_no from (select s_no,score from Sc where c_no='1') a,
(select s_no,score from Sc where c_no='2') b 
where a.score>b.score 
and a.s_no=b.s_no;


2、查询平均成绩大于60分的同学的学号和平均成绩。

select s_no,avg(score) from Sc 
group by s_no 
having avg(score)>60;


3、查询所有同学的学号、姓名、选课数、总成绩。

select Student.s_no,Student.sname,count(Sc.c_no),sum(score) 
from Student 
left outer join Sc 
on Student.s_no=Sc.s_no 
group by Student.s_no, Student.sname;


4、查询姓李的老师的个数。

select count(distinct(tname)) from Teacher where tname like '李';


5、查询没学过“叶平”老师课的同学的学号、姓名

select Student.s_no,Student.sname from Student 
where s_no not in(
select distinct (Sc.s_no) from Sc,Course,Teacher 
where Sc.s_no=Course.c_no 
and Teacher.t_no=Course.t_no 
and Teacher.tname='叶平'
);


6、查询学过“001”并且也学过编号“002”课程的同学的学号、姓名。

select Student.s_no,Student.sname from Student,Sc 
where Student.s_no=Sc.s_no 
and Sc.c_no='002' 
and exists(
select * from Sc as Sc1 
where Sc.s_no=Sc1.s_no 
and Sc1.s_no='002'
);


7、查询所有课程成绩小于60分的同学的学号、姓名。

select s_no,sname from Student 
where s_no not in (
select S.s_no from Student AS S,Sc 
where S.s_no=Sc.s_no 
and score>60
);


8、查询没有学全所有课的同学的学号、姓名。

select Student.s_no,Student.sname from Student,Sc 
where Student.s_no=Sc.s_no 
group by Student.s_no,Student.sname 
having count(c_no)<(
select count(*) from Course
);


9、查询至少学过学号为“001”同学所有一门课的其他同学学号和姓名。

select distinct s_no,sname 
from Student,Sc 
where Student.s_no=Sc.s_no 
and Sc.c_no in (
select c_no 
from Sc 
where s_no='1001'
);


10、把“sc”表中“叶平”老师教的课的成绩都更改为此课程的平均成绩。

update Sc set score=(
select avg(Sc_2.score) 
from Sc Sc_2 
where SC_2.c_no=Sc.c_no 
) 
from Course,Teacher 
where Course.c_no=Sc.c_no 
and Course.t_no=Teacher.t_no 
and Teacher.tname='叶平'
);


11、查询和“1002”号同学学习的课程完全相同的其他同学学号和姓名。

select s_no from Sc 
where c_no in (
select c_no from Sc where s_no='1002'
) 
group by s_no 
having count(*)=(
select count(*) from Sc 
where s_no='1002'
);


12、删除学习“叶平”老师课的sc表记录。

delete Sc from course,Teacher 
where Course.c_no=SC.c_no 
and Course.t_no=Teacher.t_no 
and tname='叶平';


13、向sc表中插入一些记录，这些记录要求符合一下条件：没有上过编号“003”课程的同学学号

insert into Sc 
select s_no from Student 
where s_no not in (
Select s_no from Sc where c_no='003'
);


14、查询各科成绩最高和最低的分：以如下形式显示：课程ID，最高分，最低分。

SELECT L.c_no As c_no,L.score AS max_score,R.score AS mix_score FROM Sc L ,Sc AS R

WHERE L.c_no = R.c_no and

L.score = (SELECT MAX(IL.score)

FROM Sc AS IL,Student AS IM

WHERE L.c_no = IL.c_no and IM.s_no=IL.s_no

GROUP BY IL.c_no)

AND

R.Score = (SELECT MIN(IR.score)

FROM Sc AS IR

WHERE R.c_no = IR.c_no

GROUP BY IR.c_no

) order by L.c_no;


15、查询不同老师所教不同课程平均分从高到低显示。

select c_no,avg(score) avg_score 
from Sc 
group by c_no 
order by avg_score desc 
16、统计各科成绩，各分数段人数：课程ID，课程名称，【100-85】，【85-70】，【70-60】，【<60】

select Course.c_no,cname,

count(case when score>85 and score<=100 then score end) '[85-100]',

count(case when score>70 and score<=85 then score end) '[70-85]',

count(case when score>=60 and score<=70 then score end) '[60-70]',

count(case when score<60 then score end) '[<60]'

from Course,Sc

where Course.c_no=Sc.c_no

group by Course.c_no,c_name;


17、查询每门课程被选修的学生数

select c_no,count(*) from Sc group by c_no;


18、查询出只选修了一门课程的全部学生的学号和姓名

select Student.s_no,Student.sname,count(c_no) 
from Student 
join Sc 
on Student.s_no=Sc.s_no 
group by Student.s_no, Student.sname 
having count(c_no)=1;


19、查询男生、女生人数

select count(*) from Student group by sex;


20、查询姓“张”的学生名单

select * from Student where sname like '张%';


21、查询同名同性学生名单，并统计同名人数。

select sname ,count(*) from Student group by sname having count(*)>1;


22、查询1994年出生的学生名单（注：student表中sage列的类型是datatime）

select * from Student where year(curdate())-age='1994';


23、查询每门课程的平均成绩，结果按平均成绩升序排列，平均成绩相同时，按课程号降序排列。

select c_no ,avg(score)from Sc 
group by c_no 
order by avg(score) asc,c_no desc;


24、查询平均成绩都大于85的所有学生的学号，姓名和平均成绩

select Student.s_no,Student.sname,avg(score) 
from Student,Sc 
where Student.s_no=Sc.s_no 
group by Student.s_no, Student.sname 
having avg(score)>85;


25、查询课程名称为“数据库”且分数低于60的学生姓名和分数

select Student.sname,Sc.score 
from Student,Sc 
where Student.s_no=Sc.s_no 
and Sc.score<60 
and Sc.c_no=(
select c_no from Course where cname='数据库'
);


26、查询所有学生的选课情况

select Student.s_no,Student.sname,Sc.s_no,Course.cname 
from Student,Sc,Course 
where Student.s_no=Sc.s_no 
and Sc.c_no=Course.c_no;


27、查询不及格的课程，并按课程号从大到小排序。

select Student.sname,Sc.c_no,Course.cname,Sc.score 
from Student,Sc,Course 
where Student.s_no=Sc.s_no 
and Sc.c_no=Course.c_no 
and Sc.score<60 
order by c_no;


28、查询课程编号为003且课程成绩在80分以上的学生的学号和姓名。

select Student.s_no,Student.sname 
from Student,Sc,Course 
where Sc.score>80 
and Course.c_no='003';


29、求选修了课程的学生人数。

select count(*) from (select count(*) from Sc group by s_no) b;


30、查询选修了“冯老师”所授课程的学生中，成绩最高的学生姓名及其成绩。

select Student.sname,Sc.score 
from Student,Sc,Course 
where Student.s_no=Sc.s_no 
and Sc.c_no=Course.c_no 
order by score desc 
limit 1;


31、查询各个课程及相应的选修人数。

select Course.c_no,Course.cname,count(s_no) from Course 
join Sc
on Course.c_no=Sc.c_no 
group by Course.c_no, Course.cname;


32、查询每门课程最好的前两名。

select a.s_no,a.c_no,a.score from Sc a 
where (
select count(distinct score) 
from Sc b 
where b.c_no=a.c_no 
and b.score>=a.score
)<=2 
order by a.c_no,a.score desc ;


33、查询每门课程的学生选修人数（超过10人的课程才统计）。要求输出课程号和选修人数，查询结果按人数降序排列，查询结果按人数降序排列，若人数相同，按课程号升序排列。

select Sc.c_no,count(*) from Sc 
group by c_no 
having count(*)>10 
order by count(*) desc,c_no;


34、检索至少选修两门课程的学生学号。

select s_no from Sc group by s_no having count(*)>2;


35、查询全部学生都选修的课程的课程号和课程名。

select Course.c_no,Course.cname from Course 
join Sc 
on Course.c_no=Sc.c_no 
join (
select c_no,count(s_no) from Sc 
group by c_no 
having count(s_no)=(
select count(*) from Student) )as a 
on Course.c_no=a.c_no;


36、查询两门以上不及格课程的同学的学号及其平均成绩。

select s_no,avg(score) from Sc 
where s_no in (
select s_no from Sc 
where score<60 
group by s_no 
having count(*)>2
) group by s_no;

