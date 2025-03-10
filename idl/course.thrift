namespace go course
include "model.thrift"

struct TermListRequest {}

struct TermListResponse {
    1: required model.BaseResp base
    2: required list<string> data
}

struct CourseListRequest {
    1: required string term
}

struct CourseListResponse {
    1: required model.BaseResp base
    2: required list<model.Course> data
}

struct GetCalendarRequest {
    1: required string term
}

struct GetCalendaResponse {
    1: required model.BaseResp base
    2: required string content
}

service CourseService {
    CourseListResponse GetCourseList(1: CourseListRequest req)
    TermListResponse GetTermList(1: TermListRequest req)
    GetCalendaResponse GetCalendar(1: GetCalendarRequest req)
}
