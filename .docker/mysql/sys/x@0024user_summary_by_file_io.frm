TYPE=VIEW
query=select if(isnull(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`),\'background\',`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`) AS `user`,sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`COUNT_STAR`) AS `ios`,sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`SUM_TIMER_WAIT`) AS `io_latency` from `performance_schema`.`events_waits_summary_by_user_by_event_name` where (`performance_schema`.`events_waits_summary_by_user_by_event_name`.`EVENT_NAME` like \'wait/io/file/%\') group by if(isnull(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`),\'background\',`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`) order by sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`SUM_TIMER_WAIT`) desc
md5=9ad3ed8fe2e129e434b9d87da66a32dd
updatable=0
algorithm=1
definer_user=mysql.sys
definer_host=localhost
suid=0
with_check_option=0
timestamp=2024-09-17 19:51:53
create-version=1
source=SELECT IF(user IS NULL, \'background\', user) AS user, SUM(count_star) AS ios, SUM(sum_timer_wait) AS io_latency  FROM performance_schema.events_waits_summary_by_user_by_event_name WHERE event_name LIKE \'wait/io/file/%\' GROUP BY IF(user IS NULL, \'background\', user) ORDER BY SUM(sum_timer_wait) DESC
client_cs_name=utf8
connection_cl_name=utf8_general_ci
view_body_utf8=select if(isnull(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`),\'background\',`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`) AS `user`,sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`COUNT_STAR`) AS `ios`,sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`SUM_TIMER_WAIT`) AS `io_latency` from `performance_schema`.`events_waits_summary_by_user_by_event_name` where (`performance_schema`.`events_waits_summary_by_user_by_event_name`.`EVENT_NAME` like \'wait/io/file/%\') group by if(isnull(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`),\'background\',`performance_schema`.`events_waits_summary_by_user_by_event_name`.`USER`) order by sum(`performance_schema`.`events_waits_summary_by_user_by_event_name`.`SUM_TIMER_WAIT`) desc
