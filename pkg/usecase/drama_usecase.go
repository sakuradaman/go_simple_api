package usecase

// インターフェースを定義(使用するメソッドを定義)
type Drama interface {
	FindAllDramas()
}

// 上記のインターフェースを用いて構造体を定義
type DramaUsecase struct {
}

func (du *DramaUsecase) FindAllDramas() {

}
