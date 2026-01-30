# บทที่ 5 GO Concurrency
ในบทนี้ผมเคยอ่านเรื่องนี้ก่อนจะเรียนวิชา Operating System ปรากฏว่าอ่านไม่เข้าใจเลย ดังนั้นถ้าอยากเข้าใจ GO Concurrnecy จริงๆ ผมแนะนำให้ศึกษาเรื่องนี้ โดยผมได้มีโอกาสเขียน Blog เกี่ยวกับเรื่องนี้ เลยคิดว่าน่าจะมีประโยชน์ถ้าให้ทุกท่านได้เตรียมตัวก่อนเพื่อความเข้าใจจริงๆ
- [Process](https://medium.com/@pruektanvorakul/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B-os-%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B9%80%E0%B8%97%E0%B9%88%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B8%96%E0%B8%B9%E0%B8%81-%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88-3-fd348f2d67c5)
- [Threads & Concurrency](https://medium.com/@pruektanvorakul/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B-os-%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B9%80%E0%B8%97%E0%B9%88%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B8%96%E0%B8%B9%E0%B8%81-%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88-4-47c677e97f03)
- [Synchronization Tools](https://medium.com/@pruektanvorakul/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B-os-%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B9%80%E0%B8%97%E0%B9%88%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%80%E0%B8%99%E0%B9%89%E0%B8%99%E0%B8%96%E0%B8%B9%E0%B8%81-%E0%B8%9A%E0%B8%97%E0%B8%97%E0%B8%B5%E0%B9%88-6-c284709e96b8)

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- Operating Systems 
  - พื้นฐาน Process
  - พื้นฐาน Thread
  - พื้นฐาน Race condition
  - พื้นฐาน Synchronize tools 
- ถ้าเคยเรียนรู้ fork() ในภาษา c จะทำให้เข้าใจได้ดีขึ้น แต่ไม่จำเป็น

### เนื่องจากในบทนี้การแบ่งหัวข้อเรียนแบบค่อยๆ เป็นค่อยๆ ไปจะช่วยให้ทุกท่านเข้าใจบทเรียนได้ดีที่สุดผมจึงแบ่งเป็นหลายๆ หัวข้อดังนี้

## สารบัญ
- [Tutorial 1 Goroutine](https://github.com/Nextjingjing/go-god/tree/main/05-go-concurrency/tutorial-1)