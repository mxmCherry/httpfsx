package filesystem_test

import (
	. "github.com/mxmCherry/httpfsx/internal/filesystem"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fs", func() {
	var subject *FS

	BeforeEach(func() {
		subject = New("testdata/root")
	})

	It("should return error for inexisting entities", func() {
		list, err := subject.List("/inexisting")
		Expect(err).To(HaveOccurred())
		Expect(list).To(BeNil())
	})

	It("should restrict to root dir", func() {
		list, err := subject.List("../../") // references testdata dir
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Parent.Path).To(Equal("/"))

		Expect(list.Dirs).To(HaveLen(1))
		Expect(list.Dirs[0].Path).To(Equal("/dir"))

		Expect(list.Files).To(HaveLen(1))
		Expect(list.Files[0].Path).To(Equal("/file.txt"))
	})

	It("should list root dir", func() {
		list, err := subject.List("/")
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Parent.Name).To(Equal("root"))
		Expect(list.Parent.Path).To(Equal("/"))
		Expect(list.Parent.LastMod).NotTo(BeZero())

		Expect(list.Dirs).To(HaveLen(1))
		Expect(list.Dirs[0].Name).To(Equal("dir"))
		Expect(list.Dirs[0].Path).To(Equal("/dir"))
		Expect(list.Dirs[0].LastMod).NotTo(BeZero())

		Expect(list.Files).To(HaveLen(1))
		Expect(list.Files[0].Name).To(Equal("file.txt"))
		Expect(list.Files[0].Path).To(Equal("/file.txt"))
		Expect(list.Files[0].LastMod).NotTo(BeZero())
		Expect(list.Files[0].Size).NotTo(BeZero())
	})

	It("should list dir", func() {
		list, err := subject.List("/dir")
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Parent.Name).To(Equal("dir"))
		Expect(list.Parent.Path).To(Equal("/dir"))
		Expect(list.Parent.LastMod).NotTo(BeZero())

		Expect(list.Dirs).To(HaveLen(1))
		Expect(list.Dirs[0].Name).To(Equal("subdir"))
		Expect(list.Dirs[0].Path).To(Equal("/dir/subdir"))
		Expect(list.Dirs[0].LastMod).NotTo(BeZero())

		Expect(list.Files).To(HaveLen(1))
		Expect(list.Files[0].Name).To(Equal("dir-file.txt"))
		Expect(list.Files[0].Path).To(Equal("/dir/dir-file.txt"))
		Expect(list.Files[0].LastMod).NotTo(BeZero())
		Expect(list.Files[0].Size).NotTo(BeZero())
	})

	It("should describe file", func() {
		list, err := subject.List("/dir/dir-file.txt")
		Expect(err).NotTo(HaveOccurred())

		Expect(list.Parent.Name).To(Equal("dir"))
		Expect(list.Parent.Path).To(Equal("/dir"))
		Expect(list.Parent.LastMod).NotTo(BeZero())

		Expect(list.Dirs).To(BeEmpty())

		Expect(list.Files).To(HaveLen(1))
		Expect(list.Files[0].Name).To(Equal("dir-file.txt"))
		Expect(list.Files[0].Path).To(Equal("/dir/dir-file.txt"))
		Expect(list.Files[0].LastMod).NotTo(BeZero())
		Expect(list.Files[0].Size).NotTo(BeZero())
	})

	DescribeTable("Abs",
		func(rel, expected string) {
			Expect(subject.Abs(rel)).To(Equal(expected))
		},
		Entry("simple path", "/what/ever", "testdata/root/what/ever"),
		Entry("resolve dots", "../../../what/ever/..", "testdata/root/what"),
	)

	DescribeTable("IsFile",
		func(rel string, expected bool) {
			Expect(subject.IsFile(rel)).To(Equal(expected))
		},
		Entry("file", "/file.txt", true),
		Entry("dir", "/dir", false),
		Entry("root dir", "/", false),
		Entry("inexisting", "/inexisting", false),
	)
})
