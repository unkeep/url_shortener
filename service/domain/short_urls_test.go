package domain

import (
	"context"
	"fmt"

	"github.com/gojuno/minimock/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"url_shortener/service/database"
)

var _ = Describe("ShortURLs Test", func() {
	var (
		mc          *minimock.Controller
		ctx         context.Context
		storageMock *database.URLStorageMock
		encoderMock *UrlIDEncoderMock
		u           ShortURLs
		baseURL     = "http://test.com"
	)
	BeforeEach(func() {
		mc = minimock.NewController(GinkgoT())
		ctx = context.TODO()
		storageMock = database.NewURLStorageMock(mc)
		encoderMock = NewUrlIDEncoderMock(mc)
		u = shortURLs{
			URLStorage:   storageMock,
			URLIDEncoder: encoderMock,
			BaseURL:      baseURL,
		}
	})
	AfterEach(func() {
		mc.Finish()
	})

	Context("Create", func() {
		It("should fail for invalid URL", func() {
			_, err := u.Create(ctx, "")
			Ω(err).Should(MatchError(ErrInvalidParam))
		})

		When("input URL is valid", func() {
			const inputURL = "http://example.com/asdjkasjd"
			It("should create URL in storage", func() {
				storageMock.CreateMock.Expect(ctx, inputURL).Return(0, fmt.Errorf("test"))
				_, _ = u.Create(ctx, inputURL)
			})
			When("storage fails", func() {
				BeforeEach(func() {
					storageMock.CreateMock.Return(0, fmt.Errorf("test"))
				})
				It("should fail", func() {
					_, err := u.Create(ctx, inputURL)
					Ω(err).Should(HaveOccurred())
				})
			})
			When("created in storage", func() {
				const urlID = 123
				BeforeEach(func() {
					storageMock.CreateMock.Return(urlID, nil)
				})
				It("should encode id", func() {
					encoderMock.EncodeMock.Expect(urlID).Return("", fmt.Errorf("test"))
					_, _ = u.Create(ctx, inputURL)
				})
				When("encoding fails", func() {
					BeforeEach(func() {
						encoderMock.EncodeMock.Return("", fmt.Errorf("test"))
					})
					It("should fail", func() {
						_, err := u.Create(ctx, inputURL)
						Ω(err).Should(HaveOccurred())
					})
				})
				When("encoding succeeds", func() {
					const encodedID = "abc"
					BeforeEach(func() {
						encoderMock.EncodeMock.Return(encodedID, nil)
					})
					It("should return short url", func() {
						shortURL, err := u.Create(ctx, inputURL)
						Ω(err).ShouldNot(HaveOccurred())
						wantURL := fmt.Sprintf("%s/%s", baseURL, encodedID)
						Ω(shortURL).Should(Equal(wantURL))
					})
				})
			})
		})
	})

	Context("Delete", func() {
		It("should fail for invalid URL", func() {
			err := u.Delete(ctx, "")
			Ω(err).Should(MatchError(ErrInvalidParam))
		})

		When("input URL is valid", func() {
			const shortURL = "http://example.com/qwe"
			It("should decode URL ID from path", func() {
				encoderMock.DecodeMock.Expect("qwe").Return(0, fmt.Errorf("test"))
				_ = u.Delete(ctx, shortURL)
			})
			When("decoding fails", func() {
				BeforeEach(func() {
					encoderMock.DecodeMock.Return(0, fmt.Errorf("test"))
				})
				It("should fail", func() {
					Ω(u.Delete(ctx, shortURL)).Should(MatchError(ErrInvalidParam))
				})
			})

			When("decoding succeeds", func() {
				const urlID = 321
				BeforeEach(func() {
					encoderMock.DecodeMock.Return(urlID, nil)
				})
				It("should delete URL by ID", func() {
					storageMock.DeleteByIDMock.Expect(ctx, urlID).Return(fmt.Errorf("test"))
					_ = u.Delete(ctx, shortURL)
				})
				When("storage fails", func() {
					BeforeEach(func() {
						storageMock.DeleteByIDMock.Return(fmt.Errorf("test"))
					})
					It("should fail", func() {
						Ω(u.Delete(ctx, shortURL)).ShouldNot(Succeed())
					})
				})
				When("storage succeeds", func() {
					BeforeEach(func() {
						storageMock.DeleteByIDMock.Return(nil)
					})
					It("should succeed", func() {
						Ω(u.Delete(ctx, shortURL)).Should(Succeed())
					})
				})
			})
		})
	})

	Context("DeleteByOriginURL", func() {
		It("should fail for invalid URL", func() {
			err := u.DeleteByOriginURL(ctx, "")
			Ω(err).Should(MatchError(ErrInvalidParam))
		})

		When("input URL is valid", func() {
			const originURL = "http://example.com/qwe"
			It("should delete URL", func() {
				storageMock.DeleteMock.Expect(ctx, originURL).Return(fmt.Errorf("test"))
				_ = u.DeleteByOriginURL(ctx, originURL)
			})
			When("storage fails", func() {
				BeforeEach(func() {
					storageMock.DeleteMock.Return(fmt.Errorf("test"))
				})
				It("should fail", func() {
					Ω(u.DeleteByOriginURL(ctx, originURL)).ShouldNot(Succeed())
				})
			})
			When("storage succeeds", func() {
				BeforeEach(func() {
					storageMock.DeleteMock.Return(nil)
				})
				It("should succeed", func() {
					Ω(u.DeleteByOriginURL(ctx, originURL)).Should(Succeed())
				})
			})
		})
	})

	Context("GetOriginURLByShortPath", func() {
		const shortPath = "/asd"
		It("should decode URL ID from path", func() {
			encoderMock.DecodeMock.Expect("asd").Return(0, fmt.Errorf("test"))
			_, _ = u.GetOriginURLByShortPath(ctx, shortPath)
		})
		When("decoding fails", func() {
			BeforeEach(func() {
				encoderMock.DecodeMock.Return(0, fmt.Errorf("test"))
			})
			It("should fail", func() {
				_, err := u.GetOriginURLByShortPath(ctx, shortPath)
				Ω(err).Should(MatchError(ErrInvalidParam))
			})
		})

		When("decoding succeeds", func() {
			const urlID = 222
			BeforeEach(func() {
				encoderMock.DecodeMock.Return(urlID, nil)
			})
			It("should get URL by ID", func() {
				storageMock.GetMock.Expect(ctx, urlID).Return("", fmt.Errorf("test"))
				_, _ = u.GetOriginURLByShortPath(ctx, shortPath)
			})
			When("storage fails", func() {
				BeforeEach(func() {
					storageMock.GetMock.Return("", fmt.Errorf("test"))
				})
				It("should fail", func() {
					_, err := u.GetOriginURLByShortPath(ctx, shortPath)
					Ω(err).Should(HaveOccurred())
				})
			})
			When("url not found", func() {
				BeforeEach(func() {
					storageMock.GetMock.Return("", database.ErrNotFound)
				})
				It("should fail with ErrNotFound", func() {
					_, err := u.GetOriginURLByShortPath(ctx, shortPath)
					Ω(err).Should(MatchError(ErrNotFound))
				})
			})
			When("storage succeeds", func() {
				const urlFromDB = "example.com"
				BeforeEach(func() {
					storageMock.GetMock.Return(urlFromDB, nil)
				})
				It("should return the URL from DB", func() {
					gotURL, err := u.GetOriginURLByShortPath(ctx, shortPath)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(gotURL).Should(Equal(urlFromDB))
				})
			})
		})
	})

	Context("GetByOriginURL", func() {
		It("should fail for invalid URL", func() {
			_, err := u.GetByOriginURL(ctx, "")
			Ω(err).Should(MatchError(ErrInvalidParam))
		})

		When("input URL is valid", func() {
			const originURL = "http://example.com/qwe"
			It("should find the URL", func() {
				storageMock.FindMock.Expect(ctx, originURL).Return(0, fmt.Errorf("test"))
				_, _ = u.GetByOriginURL(ctx, originURL)
			})
			When("storage fails", func() {
				BeforeEach(func() {
					storageMock.FindMock.Return(0, fmt.Errorf("test"))
				})
				It("should fail", func() {
					_, err := u.GetByOriginURL(ctx, originURL)
					Ω(err).Should(HaveOccurred())
				})
			})
			When("URL is not found", func() {
				BeforeEach(func() {
					storageMock.FindMock.Return(0, database.ErrNotFound)
				})
				It("should fail with ErrNotFound", func() {
					_, err := u.GetByOriginURL(ctx, originURL)
					Ω(err).Should(MatchError(ErrNotFound))
				})
			})
			When("storage succeeds", func() {
				const urlID = 111
				BeforeEach(func() {
					storageMock.FindMock.Return(urlID, nil)
				})
				It("should encode URL ID", func() {
					encoderMock.EncodeMock.Expect(urlID).Return("", fmt.Errorf("test"))
					_, _ = u.GetByOriginURL(ctx, originURL)
				})
				When("encoding fails", func() {
					BeforeEach(func() {
						encoderMock.EncodeMock.Return("", fmt.Errorf("test"))
					})
					It("should fail", func() {
						_, err := u.GetByOriginURL(ctx, originURL)
						Ω(err).Should(HaveOccurred())
					})
				})
				When("encoding succeeds", func() {
					const encodedID = "qwe1"
					BeforeEach(func() {
						encoderMock.EncodeMock.Return(encodedID, nil)
					})
					It("should return a short link", func() {
						gotURL, err := u.GetByOriginURL(ctx, originURL)
						Ω(err).ShouldNot(HaveOccurred())
						wantURL := fmt.Sprintf("%s/%s", baseURL, encodedID)
						Ω(gotURL).Should(Equal(wantURL))
					})
				})
			})
		})
	})
})
