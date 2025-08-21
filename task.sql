select * from departments;
select * from employees;
select * from projects;
select * from employee_projects;



select
    p.name,
    sum(e.salary) as total_salary_cost
from
    projects p
join
        employee_projects ep on p.id = ep.project_id
join
        employees e on ep.employee_id = e.id
group by
    p.name, p.id
order by
    total_salary_cost desc
limit 1;


SELECT
    p.name,
    SUM(e.salary) AS total_salary_cost
FROM
    projects p
JOIN
        employee_projects ep ON p.id = ep.project_id
JOIN
        employees e ON ep.employee_id = e.id
GROUP BY
    p.id, p.name
ORDER BY
    total_salary_cost DESC
LIMIT 1;