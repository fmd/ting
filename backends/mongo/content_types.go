package mongo

func (r *Repo) StructureType(contentType string, structure []byte) {

}

func (r *Repo) ContentTypes() ([]string, error) {
    n, err := r.Db.CollectionNames()
    if err != nil {
        return nil, err
    }

    names := make([]string,1)

    for _, name := range n {
        if name != "structures" && name != "system.indexes" {
            names = append(names, name)
        }
    }

    return names, nil
}