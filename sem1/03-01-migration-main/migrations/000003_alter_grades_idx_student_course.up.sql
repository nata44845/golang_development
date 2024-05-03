BEGIN;

CREATE INDEX student_course_idx ON Grades(StudentID,CourseID);

END;