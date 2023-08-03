//go:build integration
// +build integration

package gormrepo_test

// func TestImageRepository_Upload(t *testing.T) {
// 	t.Run("ShouldInsertImage", func(t *testing.T) {
// 		//-- init
// 		db := storage.PostgresDbConn(&dbName)
// 		defer cleanDB(t, db)

// 		fakeBinaryImage := test.FakeImageBinary(t)

// 		//-- code under test
// 		imageRepo := gormrepo.NewImageRepository(db)
// 		image, err := imageRepo.Upload(context.TODO(), fakeBinaryImage)

// 		//-- assert
// 		require.NoError(t, err)
// 		require.NotNil(t, image)
// 		require.Equal(t, fakeBinaryImage, image.BinaryImage)
// 	})
// }
