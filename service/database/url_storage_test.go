package database

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL storage", func() {
	var (
		conn = getDBConn()
		ctx  = context.Background()
		s    = NewURLStorage(conn)
		url1 = "http://url-storage-test.com/qwerty"
		url2 = "http://url-storage-test.com/12345678"
	)

	AfterEach(func() {
		Ω(s.Delete(ctx, url1)).Should(Succeed())
		Ω(s.Delete(ctx, url2)).Should(Succeed())
	})

	Context("Create", func() {
		It("should succeed and return an ID", func() {
			id, err := s.Create(ctx, url1)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(id).ShouldNot(BeZero())
		})
		When("url1 is in DB", func() {
			var id uint64
			BeforeEach(func() {
				var err error
				id, err = s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return another id on another url", func() {
				gotID, err := s.Create(ctx, url2)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(gotID > id).Should(BeTrue())
			})
			It("should return same id on subsequent calls", func() {
				gotID, err := s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(gotID).Should(Equal(id))
			})
		})
	})

	Context("Get", func() {
		It("should return NotFound err for non-existing url", func() {
			_, err := s.Get(ctx, 0)
			Ω(err).Should(MatchError(ErrNotFound))
		})
		When("url1 is in DB", func() {
			var id uint64
			BeforeEach(func() {
				var err error
				id, err = s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return the url by id", func() {
				gotURL, err := s.Get(ctx, id)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(gotURL).Should(Equal(url1))
			})
		})
	})

	Context("Find", func() {
		It("should return NotFound err for non-existing url", func() {
			_, err := s.Find(ctx, url1)
			Ω(err).Should(MatchError(ErrNotFound))
		})
		When("url1 is in DB", func() {
			var id uint64
			BeforeEach(func() {
				var err error
				id, err = s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return the url id", func() {
				gotID, err := s.Find(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(gotID).Should(Equal(id))
			})
		})
	})

	Context("DeleteByID", func() {
		It("should succeed for non-existing url", func() {
			Ω(s.DeleteByID(ctx, 0)).Should(Succeed())
		})
		When("url1 is in DB", func() {
			var id uint64
			BeforeEach(func() {
				var err error
				id, err = s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should totally delete", func() {
				Ω(s.DeleteByID(ctx, id)).Should(Succeed())
				_, err := s.Get(ctx, id)
				Ω(err).Should(MatchError(ErrNotFound))
			})
		})
	})

	Context("Delete", func() {
		It("should succeed for non-existing url", func() {
			Ω(s.Delete(ctx, url1)).Should(Succeed())
		})
		When("url1 is in DB", func() {
			var id uint64
			BeforeEach(func() {
				var err error
				id, err = s.Create(ctx, url1)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should totally delete", func() {
				Ω(s.Delete(ctx, url1)).Should(Succeed())
				_, err := s.Get(ctx, id)
				Ω(err).Should(MatchError(ErrNotFound))
			})
		})
	})
})
