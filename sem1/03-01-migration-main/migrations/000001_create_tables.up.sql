BEGIN;

CREATE TABLE Students (    
StudentID INT PRIMARY KEY,    
FirstName VARCHAR(50),    
LastName VARCHAR(50),    
EnrollmentDate DATE);

CREATE TABLE Professors (    
ProfessorID INT PRIMARY KEY,    
FirstName VARCHAR(50),    
LastName VARCHAR(50),    
Department VARCHAR(50));

CREATE TABLE Courses (    
CourseID INT PRIMARY KEY,    
CourseName VARCHAR(100),    
Department VARCHAR(50),    
Credits INT,    
ProfessorID INT,    
FOREIGN KEY (ProfessorID) REFERENCES Professors(ProfessorID));

CREATE TABLE Grades (    
GradeID INT PRIMARY KEY,    
StudentID INT,    
CourseID INT,    
Grade VARCHAR(2),    
FOREIGN KEY (StudentID) REFERENCES Students(StudentID),    
FOREIGN KEY (CourseID) REFERENCES Courses(CourseID));

END;